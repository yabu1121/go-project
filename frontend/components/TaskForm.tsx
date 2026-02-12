"use client"
import { useMutateTask } from '@/hooks/useMutateTask'
import useStore from '@/store'
import { FormEvent } from 'react'
import { PlusIcon } from '@heroicons/react/24/solid'

export const TaskForm = () => {
  const { editedTask, updateEditedTask } = useStore()
  const { createTaskMutation, updateTaskMutation } = useMutateTask()

  const handleSubmit = (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    if (editedTask.id === 0) {
      createTaskMutation.mutate({
        title: editedTask.title,
      })
    } else {
      updateTaskMutation.mutate({
        id: editedTask.id,
        title: editedTask.title,
      })
    }
  }

  return (
    <form className="mb-5 text-center" onSubmit={handleSubmit}>
      <input
        className="my-2 rounded border border-gray-300 px-3 py-2 text-sm focus:outline-none placeholder-gray-500"
        placeholder="New Task ?"
        type="text"
        onChange={(e) => updateEditedTask({ ...editedTask, title: e.target.value })}
        value={editedTask.title}
      />
      <button className="ml-2 bg-indigo-600 hover:bg-indigo-700 text-white rounded px-3 py-2 disabled:opacity-40" disabled={!editedTask.title}>
        {editedTask.id === 0 ? 'Create' : 'Update'}
      </button>
    </form>
  )
}
