async function getTask() {
  try {
    const response = await fetch("/tasks");
    const data = await response.json();
    console.log(data);

    const taskList = document.querySelector(".task-list");
    taskList.innerText = JSON.stringify(data, null, 2);
  } catch (error) {
    console.error("Error fetching task:", error);
  }
}

getTask();
