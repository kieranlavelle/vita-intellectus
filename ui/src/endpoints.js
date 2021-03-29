import { API } from './http'

function completeHabit(token, habitID){
  return API.put(`habit/${habitID}/complete`, {}, {
    headers: {
      Authorization: token,
    }
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

function editHabit(token, body, habitID){
  return API.put(`habit/${habitID}`, body, {
    headers: {
      Authorization: token,
    }
  })
}

function deleteHabit(token, habitID) {
  return API.delete(`habit/${habitID}`, {
    headers: {
      Authorization: token,
    }
  })
}

export {
  completeHabit,
  getHabits,
  createHabit,
  editHabit,
  deleteHabit
}