import axios from 'axios'

// apiインスタンスを作成
export const api = axios.create({
  baseURL: process.env.NEXT_PUBLIC_API_URL,
  withCredentials: true,
})

// csrfトークンを取得
export const csrf = async () => {
  const { data } = await api.get('/csrf-token')
  api.defaults.headers.common['X-CSRF-Token'] = data.csrf_token
}
