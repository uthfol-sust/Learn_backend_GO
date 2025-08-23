import {renderTask ,createTaskListHeader ,changeStatus } from "./utils_func.js"
import { setupModal } from "./add_task_render.js";
import { getTasks, createTask } from "./handle_API/handler.js";


const taskList = document.querySelector(".task-list");

const icons_object = [
  "fa-solid fa-circle-check" , 
  "fa-regular fa-circle",  
];

// add header (add/search/filter)
taskList.append(createTaskListHeader());

// fetch tasks
getTasks(taskList, icons_object);

// setup modal logic
setupModal(taskList);

