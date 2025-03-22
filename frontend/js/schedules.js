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

    function parseJwt(token) {
        try {
            const base64Url = token.split(".")[1]; // Берем payload
            const base64 = base64Url.replace(/-/g, "+").replace(/_/g, "/");
            const jsonPayload = decodeURIComponent(
                atob(base64)
                    .split("")
                    .map((c) => "%" + ("00" + c.charCodeAt(0).toString(16)).slice(-2))
                    .join("")
            );
            return JSON.parse(jsonPayload);
        } catch (e) {
            return null;
        }
    }

// Получаем информацию о пользователе
    const token = localStorage.getItem("jwtToken");
    const user = token ? parseJwt(token) : null;
    const isChef = user?.isChef || false; // Проверяем, шеф ли он

    if (!isChef) {
        saveChangesBtn.style.display = 'none';
    }

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
    const editBtn = document.getElementById("schedule_edit-schedule-btn")
    if (!isChef) {
        editBtn.style.display = 'none';
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

    const delBtn = document.getElementById("schedule_delete-schedule-btn")
    if (!isChef) {
        delBtn.style.display = 'none';
    }
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


    // Получаем ссылки на элементы модального окна
    const createEventModal = document.getElementById("schedule_create-event-modal");
    const createEventForm = document.getElementById("schedule_create-event-form");
    const createEventBtn = document.getElementById("schedule_create-event-btn");
    const createCloseBtn = document.getElementById("schedule_create-close");

// Поля ввода
    const eventNameInput = document.getElementById("schedule_event-name");
    const eventDateInput = document.getElementById("schedule_event-date");
    const eventLocationInput = document.getElementById("schedule_event-location");

// Открытие модального окна
    createEventBtn.addEventListener("click", () => {
        createEventModal.style.display = "flex";
    });

// Закрытие модального окна
    createCloseBtn.addEventListener("click", () => {
        createEventModal.style.display = "none";
    });


// Закрытие при клике вне формы
    window.addEventListener("click", (event) => {
        if (event.target === createEventModal) {
            createEventModal.style.display = "none";
        }
    });

// Обработчик отправки формы
    createEventForm.addEventListener("submit", (e) => {
        e.preventDefault();

        const formattedDate = new Date(eventDateInput.value).toISOString();

        const newEvent = {
            event_name: eventNameInput.value,
            event_date: formattedDate,
            location: eventLocationInput.value
        };

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
                .then(() => {
                    alert("Событие успешно создано!");
                    createEventModal.style.display = "none"; // Закрыть модальное окно
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
