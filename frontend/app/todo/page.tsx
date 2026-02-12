"use client"
import { TaskList } from '@/components/TaskList'
import { TaskForm } from '@/components/TaskForm'
import { ArrowLeftOnRectangleIcon } from '@heroicons/react/24/solid'
import { useAuth } from '@/hooks/useAuth'

export default function Todo() {
  const { logoutMutation } = useAuth()

  return (
    <div className="flex flex-col items-center justify-center min-h-screen text-gray-600 bg-gray-100 font-mono">
      <div className="flex items-center mb-5">
        <ArrowLeftOnRectangleIcon
          className="h-6 w-6 cursor-pointer text-blue-500 mr-2"
          onClick={() => logoutMutation.mutate()}
        />
        <span className="text-xl font-bold">Log out</span>
      </div>
      <TaskForm />
      <TaskList />
    </div>
  )
}
