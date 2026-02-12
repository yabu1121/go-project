export type Task = {
  id: number
  title: string
  created_at: string
  updated_at: string
}

export type CsrfToken = {
  csrf_token: string
}

export type Credentials = {
  email: string
  password: string
}