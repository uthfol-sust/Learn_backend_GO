import { getTasks, createTask  , deleteTask} from "./handle_API/handler.js";

// tasks.js
// export const tasks = [
//   {
//     id: 1,
//     title: "Finish Golang project",
//     description: "Complete the REST API using Gorilla Mux",
//     status: "pending",
//     dueDate: new Date(Date.now() + 1 * 60 * 60 * 1000),
//     priority: "high",
//     category: "work",
//     createdAt: new Date(),
//     updatedAt: new Date(),
//     notifications: true
//   },
//   {
//     id: 2,
//     title: "Study JavaScript",
//     description: "Review ES6+ features and async programming",
//     status: "in-progress",
//     dueDate: new Date(Date.now() + 2 * 24 * 60 * 60 * 1000),
//     priority: "medium",
//     category: "study",
//     createdAt: new Date(),
//     updatedAt: new Date(),
//     notifications: false
//   },
//   {
//     id: 3,
//     title: "Buy groceries",
//     description: "Milk, eggs, vegetables",
//     status: "completed",
//     dueDate: new Date(Date.now() + 6 * 60 * 60 * 1000), // 6 hours later
//     priority: "low",
//     category: "personal",
//     createdAt: new Date(),
//     updatedAt: new Date(),
//     notifications: true
//   }
// ];

// render a single task card
export function renderTask(task, container) {
  const div = document.createElement("div");
  div.classList.add("task-card");

  const title = document.createElement("h4");
  title.textContent = task.title;

  const status = document.createElement("div");
  status.classList.add("status");

  const icon = document.createElement("i");
  icon.className = "fa-regular fa-circle status-icon";

  const textNode = document.createElement("p");
  textNode.textContent = task.status;
  textNode.classList.add("status-text");

  if (task.status === "completed") {
    icon.style.color = "green";
    textNode.style.color = "green";
  }
  status.append(icon, textNode);

  const deleteicon = document.createElement("i");
  deleteicon.className = "fa-solid fa-trash delete-task"

  // ✅ attach listener here so you get the right id
  deleteicon.addEventListener("click", () => {
    const ok = confirm("Are you sure?");
    if (ok) {
      console.log("Deleting task id:", task.id); 
      deleteTask(task.id); // <-- call your API here
      div.remove();        // remove from UI
    }
  });


  div.append(title, status,deleteicon);

  container.append(div);

}

// ui.js
export function createTaskListHeader() {
  const header = document.createElement("div");
  header.classList.add("taskListHeader");

  // Add Task button
  const addTask = document.createElement("div")
  addTask.classList.add("addTask")

  const addTaskIcon = document.createElement("i")
  addTaskIcon.classList.add("fa-solid", "fa-plus")

  const addTaskText = document.createElement("p")
  addTaskText.textContent = "Add"

  addTask.append(addTaskIcon, addTaskText)

  addTask.addEventListener("click", (e)=>{
    document.querySelector("#taskModal").classList.add("active");
  })

  // Search Task
  const searchTask = document.createElement("div");
  searchTask.classList.add("searchTask");

  const searchInput = document.createElement("input");
  searchInput.type = "text";
  searchInput.placeholder = "Search tasks...";

  const searchIcon = document.createElement("i");
  searchIcon.classList.add("fa-solid", "fa-magnifying-glass");

  searchTask.append(searchIcon, searchInput);

  // Filter Task
  const filterTask = document.createElement("div");
  filterTask.classList.add("filterTask");

  const filterTaskIcon = document.createElement("i");
  filterTaskIcon.classList.add("fa-solid", "fa-filter");

  const filterTaskText = document.createElement("p");
  filterTaskText.textContent = "Filter";

  filterTask.append(filterTaskIcon ,filterTaskText);

  header.append(addTask, searchTask, filterTask);
  
  return header;
}

export function changeStatus(icons_object){
  const statusIcons = document.querySelectorAll(".status-icon");

  statusIcons.forEach(icon => {
    icon.addEventListener("click", (e) => {
    const el = e.target;

    const iconClasses = el.className
      .split(" ")
      .filter(c => c !== "fa" && c !== "status-icon")
      .join(" ");

    let currentIndex = icons_object.findIndex(cls => cls === iconClasses);
    let nextIndex = (currentIndex + 1) % icons_object.length;

    el.className = "fa " + icons_object[nextIndex] + " status-icon";
    
  });
});
}

//Delete task 
function delete_Task(task){
  const deleteTasks = document.querySelectorAll(".delete-task")

  deleteTasks.forEach(item =>{
      item.addEventListener("click",(e)=>{
        const ok = confirm("Are you sure?")
        if(ok){
          console.log(task.id)
          
        }
      })
  })
}

