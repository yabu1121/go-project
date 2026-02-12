import { TrashIcon, PencilSquareIcon } from '@heroicons/react/24/solid'
import { Task } from '@/types'
import useStore from '@/store'
import { useMutateTask } from '@/hooks/useMutateTask'
import { memo } from 'react'

export const TaskItem = memo(({ id, title }: Omit<Task, 'created_at' | 'updated_at'>) => {
  const updateEditedTask = useStore((state) => state.updateEditedTask)
  const { deleteTaskMutation } = useMutateTask()

  return (
    <li className="my-3 text-lg font-extrabold flex justify-between">
      <span>{title}</span>
      <div className="flex float-right ml-20">
        <PencilSquareIcon
          className="h-5 w-5 mx-1 text-blue-500 cursor-pointer"
          onClick={() => {
            updateEditedTask({ id, title })
          }}
        />
        <TrashIcon
          className="h-5 w-5 mx-1 text-blue-500 cursor-pointer"
          onClick={() => {
            deleteTaskMutation.mutate(id)
          }}
        />
      </div>
    </li>
  )
})

TaskItem.displayName = 'TaskItem'
