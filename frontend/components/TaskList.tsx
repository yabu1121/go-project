import { useQueryTasks } from '@/hooks/useQueryTasks'
import { TaskItem } from './TaskItem'

export const TaskList = () => {
  const { data: tasks, status } = useQueryTasks()

  if (status === 'pending') return <div className="my-5 text-center">Loading...</div>
  if (status === 'error') return <div className="my-5 text-center text-red-500">Error</div>

  return (
    <ul className="my-5 list-none">
      {tasks?.map((task) => (
        <TaskItem key={task.id} id={task.id} title={task.title} />
      ))}
    </ul>
  )
}
