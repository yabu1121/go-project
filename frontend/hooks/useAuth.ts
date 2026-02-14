import { useMutation } from '@tanstack/react-query'
import { useRouter } from 'next/navigation'
import { api, csrf } from '@/lib/axios'
import { Credentials } from '@/types'

export const useAuth = () => {
  const router = useRouter()

  // Login関数
  const loginMutation = useMutation({
    mutationFn: async (user: Credentials) => {
      await csrf()
      return await api.post('/login', user)
    },
    onSuccess: () => {
      router.push('/todo')
    },
    onError: (err: any) => {
      // Handle error
    },
  })

  // Register関数
  const registerMutation = useMutation({
    mutationFn: async (user: Credentials) => {
      await csrf()
      return await api.post('/signup', user)
    },
    onSuccess: () => {
      router.push('/todo')
    },
  })

  // Logout関数
  const logoutMutation = useMutation({
    mutationFn: async () => {
      await csrf()
      return await api.post('/logout')
    },
    onSuccess: () => {
      router.refresh()
      router.push('/')
    },
  })

  return { loginMutation, registerMutation, logoutMutation }
}
