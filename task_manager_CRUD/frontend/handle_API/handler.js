import { renderTask , changeStatus } from "../utils_func.js";

const API ="http://localhost:8000/tasks"

export const getTasks = async (taskList, icons_object) => {
  try {
    const response = await fetch(API);
    const data = await response.json();

    data.forEach(task => {
      renderTask(task, taskList);
    });

    changeStatus(icons_object);

  } catch (error) {
    console.error("Error fetching tasks:", error);
  }
};

export const createTask = async (newTask) => {
  const response = await fetch(API, {
    method:"POST",
    headers:{ "Content-Type":"application/json" },
    body: JSON.stringify(newTask)
  });

  const data = await response.json()
  console.log("Created Task:", data)
};

export const deleteTask = async (id) => {
  try {
    const response = await fetch(`${API}/${id}`, {
      method: "DELETE",
      headers: {
        "Content-Type": "application/json"
      }
    });

    if (!response.ok) {
      throw new Error(`Failed to delete task with ID ${id}, status: ${response.status}`);
    }
    
    let data = null;
    if (response.status !== 204) {
      data = await response.json();
    }

    console.log("Deleted Task:", data ?? `Task ${id} deleted successfully`);

  } catch (error) {
    console.error("Error deleting task:", error);
  }
};
