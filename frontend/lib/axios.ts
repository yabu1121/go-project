import axios from 'axios'

export const api = axios.create({
  baseURL: process.env.NEXT_PUBLIC_API_URL,
  withCredentials: true,
})

export const csrf = async () => {
  const { data } = await api.get('/csrf-token')
  api.defaults.headers.common['X-CSRF-Token'] = data.csrf_token
}
