document.addEventListener("DOMContentLoaded", () => {
    const schedulesContainer = document.getElementById("schedule_schedules-container");
    const modal = document.getElementById("schedule_schedule-modal");
    const closeModal = document.querySelector(".schedule_close");
    const editForm = document.getElementById("schedule_edit-form");
    const saveChangesBtn = document.getElementById("schedule_save-changes");

    const scheduleName = document.getElementById("schedule_schedule-name");
    const scheduleDate = document.getElementById("schedule_schedule-date");
    const scheduleLocation = document.getElementById("schedule_schedule-location");

    const editName = document.getElementById("schedule_edit-name");
    const editDate = document.getElementById("schedule_edit-date");
    const editLocation = document.getElementById("schedule_edit-location");

    const createEventBtn = document.getElementById("schedule_create-event-btn");
    const token = localStorage.getItem("jwtToken");
    let currentScheduleId = null;

    // Загрузка всех событий
    function fetchSchedules() {
        fetch("http://localhost:8080/api/schedules/all/", {
            headers: { "Authorization": `Bearer ${token}` }
        })
            .then(response => response.json())
            .then(data => {
                schedulesContainer.innerHTML = "";
                data.forEach(schedule => {
                    const card = document.createElement("div");
                    card.className = "schedule_schedule-card";
                    card.innerHTML = `
                    <h3>${schedule.event_name}</h3>
                    <p><strong>Дата:</strong> ${new Date(schedule.event_date).toLocaleDateString()}</p>
                    <p><strong>Место:</strong> ${schedule.location}</p>
                    <button class="schedule_details-btn" data-id="${schedule.id}">Подробнее</button>
                `;
                    schedulesContainer.appendChild(card);
                });

                // Обработчик кнопок "Подробнее"
                document.querySelectorAll(".schedule_details-btn").forEach(button => {
                    button.addEventListener("click", () => {
                        const scheduleId = button.getAttribute("data-id");
                        fetchScheduleDetails(scheduleId);
                    });
                });
            })
            .catch(error => console.error("Ошибка загрузки событий:", error));
    }

    // Загрузка деталей события
    function fetchScheduleDetails(id) {
        fetch(`http://localhost:8080/api/schedules/${id}`, {
            headers: { "Authorization": `Bearer ${token}` }
        })
            .then(response => response.json())
            .then(schedule => {
                currentScheduleId = id;
                scheduleName.textContent = schedule.event_name;
                scheduleDate.textContent = new Date(schedule.event_date).toLocaleDateString();
                scheduleLocation.textContent = schedule.location;

                // Заполняем форму редактирования
                editName.value = schedule.event_name;
                editDate.value = new Date(schedule.event_date).toISOString().split('T')[0];
                editLocation.value = schedule.location;

                modal.style.display = "flex";
            })
            .catch(error => console.error("Ошибка загрузки события:", error));
    }

    // Открытие формы редактирования
    document.getElementById("schedule_edit-schedule-btn").addEventListener("click", () => {
        editForm.style.display = "block";
    });

    saveChangesBtn.addEventListener("click", () => {
        // Преобразуем дату в нужный формат с временем
        const formattedDate = new Date(editDate.value).toISOString().split('T')[0] + "T00:00:00Z";

        const updatedSchedule = {
            event_name: editName.value,
            event_date: formattedDate, // Используем отформатированную дату
            location: editLocation.value
        };

        fetch(`http://localhost:8080/api/schedules/${currentScheduleId}`, {
            method: "PUT",
            headers: {
                "Content-Type": "application/json",
                "Authorization": `Bearer ${token}`
            },
            body: JSON.stringify(updatedSchedule)
        })
            .then(response => {
                if (!response.ok) throw new Error("Ошибка обновления");
                return response.json();
            })
            .then(() => {
                alert("Событие обновлено!");
                modal.style.display = "none";
                fetchSchedules();
            })
            .catch(error => console.error("Ошибка обновления события:", error));
    });

    // Удаление события
    document.getElementById("schedule_delete-schedule-btn").addEventListener("click", () => {
        if (!confirm("Вы уверены, что хотите удалить это событие?")) return;

        fetch(`http://localhost:8080/api/schedules/${currentScheduleId}`, {
            method: "DELETE",
            headers: { "Authorization": `Bearer ${token}` }
        })
            .then(response => {
                if (!response.ok) throw new Error("Ошибка удаления");
                return response.json();
            })
            .then(() => {
                alert("Событие удалено!");
                modal.style.display = "none";
                fetchSchedules();
            })
            .catch(error => console.error("Ошибка удаления события:", error));
    });

    // Закрытие модального окна
    closeModal.addEventListener("click", () => {
        modal.style.display = "none";
        editForm.style.display = "none";
    });

    window.addEventListener("click", (event) => {
        if (event.target === modal) {
            modal.style.display = "none";
            editForm.style.display = "none";
        }
    });

    // Создание нового события
    const createEventFormContainer = document.getElementById("create-event-form-container");
    const createEventForm = document.getElementById("create-event-form");
    const cancelCreateEventBtn = document.getElementById("cancel-create-event-btn");

    const eventNameInput = document.getElementById("event-name");
    const eventDateInput = document.getElementById("event-date");
    const eventLocationInput = document.getElementById("event-location");

    // Открыть форму создания события
    createEventBtn.addEventListener("click", () => {
        createEventFormContainer.style.display = "block";
    });

    // Отменить создание события и скрыть форму
    cancelCreateEventBtn.addEventListener("click", () => {
        createEventFormContainer.style.display = "none";
    });

    // Обработчик отправки формы
    createEventForm.addEventListener("submit", (e) => {
        e.preventDefault();

        const newEvent = {
            event_name: eventNameInput.value,
            event_date: eventDateInput.value,
            location: eventLocationInput.value
        };

        // Проверка наличия всех данных
        if (newEvent.event_name && newEvent.event_date && newEvent.location) {
            fetch("http://localhost:8080/api/schedules", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    "Authorization": `Bearer ${token}`
                },
                body: JSON.stringify(newEvent)
            })
                .then(response => response.json())
                .then(data => {
                    alert("Событие успешно создано!");
                    createEventFormContainer.style.display = "none"; // Закрыть форму
                    fetchSchedules(); // Обновить список событий
                })
                .catch(error => {
                    console.error("Ошибка создания события:", error);
                    alert("Не удалось создать событие.");
                });
        } else {
            alert("Пожалуйста, заполните все поля!");
        }
    });

    fetchSchedules(); // Загружаем события при загрузке страницы
});
