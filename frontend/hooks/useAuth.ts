import { useMutation } from '@tanstack/react-query'
import { useRouter } from 'next/navigation'
import { api, csrf } from '@/lib/axios'
import { Credentials } from '@/types'

export const useAuth = () => {
  const router = useRouter()

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

  const registerMutation = useMutation({
    mutationFn: async (user: Credentials) => {
      await csrf()
      return await api.post('/signup', user)
    },
    onSuccess: () => {
      router.push('/todo')
    },
  })

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
