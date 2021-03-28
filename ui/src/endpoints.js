import { API } from './http'

function completeHabit(habitID){
  API.put(`habit/${habitID}/complete`)
    .then(response => {
      return response;
    })
    .catch(error => {
      return error.response;
    })
}

function getHabits(token){
  return API.get('habits', {
    headers: {
      Authorization: token,
    }
  })
}

function createHabit(token, body){
  return API.post('habit', body, {
    headers: {
      Authorization: token,
    }
  })
}

export {
  completeHabit,
  getHabits,
  createHabit
}