import { API } from './http'

function completeTask(token, task_id){
  return API.put(`task/complete/${task_id}`, {}, {
    headers: {
      Authorization: token,
    }
  })
}

function uncompleteTask(token, task_id){
  return API.put(`task/uncomplete/${task_id}`, {}, {
    headers: {
      Authorization: token,
    }
  })
}

function getTasks(token, filter){
  return API.get(`tasks?filter=${filter}`, {
    headers: {
      Authorization: token,
    }
  })
}

function createTask(token, body){
  return API.post('task', body, {
    headers: {
      Authorization: token,
    }
  })
}

function editHabit(token, body, habitID){
  return API.put(`habit/${habitID}`, body, {
    headers: {
      Authorization: token,
    }
  })
}

function deleteTask(token, taskID) {
  return API.delete(`task/${habitID}`, {
    headers: {
      Authorization: token,
    }
  })
}

export {
  completeTask,
  getTasks,
  uncompleteTask,
  createTask,
  editHabit,
  deleteTask
}