import { useMutation, useQueryClient } from '@tanstack/react-query'
import { useRouter } from 'next/navigation'
import { api } from '@/lib/axios'
import useStore from '@/store'
import { Task } from '@/types'

export const useMutateTask = () => {
  const queryClient = useQueryClient()
  const router = useRouter()
  const resetEditedTask = useStore((state) => state.resetEditedTask)

  const createTaskMutation = useMutation({
    mutationFn: async (task: Omit<Task, 'id' | 'created_at' | 'updated_at'>) =>
      (await api.post<Task>('/tasks', task)).data,
    onSuccess: (res) => {
      const previousTasks = queryClient.getQueryData<Task[]>(['tasks'])
      if (previousTasks) {
        queryClient.setQueryData(['tasks'], [...previousTasks, res])
      }
      resetEditedTask()
    },
    onError: (err: any) => {
      if (err.response.status === 401 || err.response.status === 403) {
        router.push('/')
        resetEditedTask()
      }
    },
  })

  const updateTaskMutation = useMutation({
    mutationFn: async (task: Omit<Task, 'created_at' | 'updated_at'>) =>
      (await api.put<Task>(`/tasks/${task.id}`, task)).data,
    onSuccess: (res, variables) => {
      const previousTasks = queryClient.getQueryData<Task[]>(['tasks'])
      if (previousTasks) {
        queryClient.setQueryData<Task[]>(
          ['tasks'],
          previousTasks.map((task) => (task.id === variables.id ? res : task))
        )
      }
      resetEditedTask()
    },
    onError: (err: any) => {
      if (err.response.status === 401 || err.response.status === 403) {
        router.push('/')
        resetEditedTask()
      }
    },
  })

  const deleteTaskMutation = useMutation({
    mutationFn: async (id: number) =>
      await api.delete(`/tasks/${id}`),
    onSuccess: (_, variables) => {
      const previousTasks = queryClient.getQueryData<Task[]>(['tasks'])
      if (previousTasks) {
        queryClient.setQueryData<Task[]>(
          ['tasks'],
          previousTasks.filter((task) => task.id !== variables)
        )
      }
      resetEditedTask()
    },
    onError: (err: any) => {
      if (err.response.status === 401 || err.response.status === 403) {
        router.push('/')
        resetEditedTask()
      }
    },
  })

  return { createTaskMutation, updateTaskMutation, deleteTaskMutation }
}
