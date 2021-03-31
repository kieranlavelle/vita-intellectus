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

function habitInfo(token, habitID) {
  return API.get(`habit/${habitID}/info`, {
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

function getFullHabits(token) {
  API.get('habits', {
    headers: {
      Authorization: token,
    }
  }).then(response => {
    let habits = response.data.habits;
    habits = habits.map((habit, index) => {
      habitInfo(token, habit.id)
        .then(response => {
          return Object.assign({}, habit, response.data.info);
        })
        .catch(error => {
          console.log(error);
        })
    })
  })
}

export {
  completeHabit,
  getHabits,
  habitInfo,
  createHabit,
  editHabit,
  deleteHabit
}