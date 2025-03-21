document.addEventListener("DOMContentLoaded", () => {
    const recipesContainer = document.getElementById("recipes-container");
    const modal = document.getElementById("recipe-modal");
    const closeModal = document.querySelector(".close");
    const createRecipeModal = document.getElementById("create-recipe-modal");
    const createRecipeButton = document.getElementById("create-recipe-button");
    const createRecipeForm = document.getElementById("create-recipe-form");

    const recipeTitle = document.getElementById("recipe-title");
    const recipeIngredients = document.getElementById("recipe-ingredients");
    const recipeText = document.getElementById("recipe-text");

    const editForm = document.getElementById("edit-form");
    const editTitle = document.getElementById("edit-title");
    const editIngredients = document.getElementById("edit-ingredients");
    const editRecipeText = document.getElementById("edit-recipe-text");
    const saveChangesBtn = document.getElementById("save-changes");

    const newRecipeTitle = document.getElementById("new-recipe-title");
    const newRecipeIngredients = document.getElementById("new-recipe-ingredients");
    const newRecipeText = document.getElementById("new-recipe-text");

    const editButton = document.getElementById("edit-button");
    const deleteButton = document.getElementById("delete-button");

    const token = localStorage.getItem("jwtToken");
    let currentRecipeId = null;

    // Функция загрузки всех рецептов
    function fetchRecipes() {
        fetch("http://localhost:8080/api/recipes/all/", {
            headers: { "Authorization": `Bearer ${token}` }
        })
            .then(response => response.json())
            .then(data => {
                recipesContainer.innerHTML = "";
                data.forEach(recipe => {
                    const card = document.createElement("div");
                    card.className = "recipe-card";
                    card.innerHTML = `
                    <h3>${recipe.title}</h3>
                    <p><strong>Ингредиенты:</strong> ${Object.keys(recipe.ingredients).length} шт.</p>
                    <button class="details-btn" data-id="${recipe.id}">Подробнее</button>
                `;
                    recipesContainer.appendChild(card);
                });

                document.querySelectorAll(".details-btn").forEach(button => {
                    button.addEventListener("click", () => {
                        const recipeId = button.getAttribute("data-id");
                        fetchRecipeDetails(recipeId);
                    });
                });
            })
            .catch(error => console.error("Ошибка загрузки рецептов:", error));
    }

    // Функция загрузки конкретного рецепта
    function fetchRecipeDetails(id) {
        fetch(`http://localhost:8080/api/recipes/${id}`, {
            headers: { "Authorization": `Bearer ${token}` }
        })
            .then(response => response.json())
            .then(recipe => {
                currentRecipeId = id;

                recipeTitle.textContent = recipe.title;
                recipeIngredients.innerHTML = Object.entries(recipe.ingredients)
                    .map(([ingredient, quantity]) => `<li>${ingredient}: ${quantity} г</li>`)
                    .join("");
                recipeText.textContent = recipe.recipe_text;

                // Заполняем форму редактирования
                editTitle.value = recipe.title;
                editIngredients.value = JSON.stringify(recipe.ingredients, null, 2);
                editRecipeText.value = recipe.recipe_text;

                modal.style.display = "flex"; // Показываем модальное окно
            })
            .catch(error => console.error("Ошибка загрузки рецепта:", error));
    }

    // Кнопка "Редактировать"
    editButton.addEventListener("click", () => {
        editForm.style.display = "block"; // Показываем форму редактирования
    });

    // Кнопка "Сохранить изменения"
    saveChangesBtn.addEventListener("click", () => {
        const updatedRecipe = {
            title: editTitle.value,
            ingredients: JSON.parse(editIngredients.value),
            recipe_text: editRecipeText.value
        };

        fetch(`http://localhost:8080/api/recipes/${currentRecipeId}`, {
            method: "PUT",
            headers: {
                "Content-Type": "application/json",
                "Authorization": `Bearer ${token}`
            },
            body: JSON.stringify(updatedRecipe)
        })
            .then(response => {
                if (!response.ok) throw new Error("Ошибка обновления");
                return response.json();
            })
            .then(() => {
                alert("Рецепт обновлен!");
                modal.style.display = "none";
                fetchRecipes(); // Перезагружаем рецепты
            })
            .catch(error => console.error("Ошибка обновления рецепта:", error));
    });

    // Кнопка "Удалить"
    deleteButton.addEventListener("click", () => {
        if (!confirm("Вы уверены, что хотите удалить рецепт?")) return;

        fetch(`http://localhost:8080/api/recipes/${currentRecipeId}`, {
            method: "DELETE",
            headers: { "Authorization": `Bearer ${token}` }
        })
            .then(response => {
                if (!response.ok) throw new Error("Ошибка удаления");
                return response.json();
            })
            .then(() => {
                alert("Рецепт удален!");
                modal.style.display = "none";
                fetchRecipes(); // Перезагружаем рецепты
            })
            .catch(error => console.error("Ошибка удаления рецепта:", error));
    });

    // Закрытие модального окна
    closeModal.addEventListener("click", () => {
        modal.style.display = "none";
        editForm.style.display = "none"; // Скрываем форму редактирования
    });

    window.addEventListener("click", (event) => {
        if (event.target === modal) {
            modal.style.display = "none";
            editForm.style.display = "none"; // Скрываем форму редактирования
        }
    });
    // Открытие модального окна для создания рецепта
    createRecipeButton.addEventListener("click", () => {
        createRecipeModal.style.display = "flex";
    });

    // Закрытие модального окна
    closeModal.addEventListener("click", () => {
        createRecipeModal.style.display = "none";
    });

    // Функция создания нового рецепта
    createRecipeForm.addEventListener("submit", (e) => {
        e.preventDefault();

        const newRecipe = {
            title: newRecipeTitle.value,
            ingredients: JSON.parse(newRecipeIngredients.value),
            recipe_text: newRecipeText.value,
        };

        fetch("http://localhost:8080/api/recipes", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                "Authorization": `Bearer ${token}`,
            },
            body: JSON.stringify(newRecipe),
        })
            .then(response => response.json())
            .then(data => {
                if (data) {
                    alert("Рецепт успешно создан!");
                    createRecipeModal.style.display = "none"; // Закрываем модальное окно
                    fetchRecipes(); // Перезагружаем рецепты
                }
            })
            .catch(error => {
                console.error("Ошибка создания рецепта:", error);
                alert("Не удалось создать рецепт.");
            });
    });


    fetchRecipes(); // Загружаем рецепты при загрузке страницы
});
