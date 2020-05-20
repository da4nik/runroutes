const baseURL = 'http://localhost:3333/api/v1'

const request = (method, path, body) => {
  return fetch(baseURL + path, {
    method: method,
    body: JSON.stringify(body),
    headers: {
      'Content-Type': 'application/json; charset=utf-8'
    }
  }).then(response => {
    return {
      status: response.status,
      data: response.json()
    }
  })
}

const points = {
  list () {
    return request('GET', '/points')
  },

  create (point) {
    return request('POST', '/points', point)
  }
}

export default {
  points
}
