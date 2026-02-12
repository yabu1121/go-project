import { useQuery } from '@tanstack/react-query'
import { api } from '@/lib/axios'
import { Task } from '@/types'

export const useQueryTasks = () => {
  return useQuery<Task[], Error>({
    queryKey: ['tasks'],
    queryFn: async () => {
      const { data } = await api.get<Task[]>('/tasks')
      return data
    },
  })
}
