import { renderTask } from "./utils_func.js"
import { getTasks,createTask } from "./handle_API/handler.js"


export function setupModal(taskList) {
  const taskModal = document.getElementById("taskModal")
  const saveTask = document.getElementById("saveTask")
  const closeModal = document.getElementById("closeModal")

  saveTask.addEventListener("click", () => {
    const name = document.getElementById("taskName").value
    const status = document.getElementById("taskStatus").value
    const due = document.getElementById("taskDue").value
    const desc = document.getElementById("taskDesc").value

    if (!name.trim()) {
      alert("Task name required!");
      return;
    }

    const newTaskObj = {
      id: Date.now(),
      title: name,
      description: desc,
      status: status,
      dueDate: due,
      priority: "medium",
      category: "general",
      createdAt: new Date(),
      updatedAt: new Date(),
      notifications: false
    };

    createTask(newTaskObj)
    renderTask(newTaskObj,taskList)


    taskModal.classList.remove("active")
  });

  closeModal.addEventListener("click", () => {
    taskModal.classList.remove("active")
  });
}
