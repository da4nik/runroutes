const baseURL = 'http://localhost:3333/api/v1'

const request = (method, path, body) => {
  let status = 200

  return fetch(baseURL + path, {
    method: method,
    body: JSON.stringify(body),
    headers: {
      'Content-Type': 'application/json; charset=utf-8'
    }
  }).then(response => {
    status = response.status
    return response.json()
  }).then(json => {
    return {
      status: status,
      data: json
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

const ways = {
  list () {
    return request('GET', '/ways')
  },

  create (way) {
    return request('POST', '/ways', way)
  }
}

export default {
  points, ways
}
