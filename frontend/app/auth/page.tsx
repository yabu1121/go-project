"use client"
import { useAuth } from '@/hooks/useAuth'
import { FormEvent, useState } from 'react'
import { CheckBadgeIcon, ArrowRightOnRectangleIcon } from '@heroicons/react/24/solid'
import { useRouter } from 'next/navigation'

export default function Auth() {
  const { loginMutation, registerMutation } = useAuth()
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const [isLogin, setIsLogin] = useState(true)
  const router = useRouter()

  const handleSubmit = async (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    if (isLogin) {
      loginMutation.mutate({
        email,
        password,
      })
      router.push('/todo')
    } else {
      registerMutation.mutate({
        email,
        password,
      })
      router.push('/todo')
    }
  }

  return (
    <div className="flex justify-center items-center flex-col min-h-screen text-gray-600 font-mono">
      <div className="flex items-center">
        <CheckBadgeIcon className="h-8 w-8 mr-2 text-blue-500" />
        <span className="text-center text-3xl font-extrabold">
          Todo App by Golang/Next.js
        </span>
      </div>
      <h2 className="my-6">{isLogin ? 'Login' : 'Create a new account'}</h2>
      <form onSubmit={handleSubmit}>
        <div>
          <input
            name="email"
            type="email"
            autoComplete="email"
            required
            className="my-2 rounded border border-gray-300 px-3 py-2 text-sm focus:outline-none"
            placeholder="Email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
          />
        </div>
        <div>
          <input
            name="password"
            type="password"
            autoComplete="current-password"
            required
            className="my-2 rounded border border-gray-300 px-3 py-2 text-sm focus:outline-none"
            placeholder="Password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
          />
        </div>
        <div className="flex justify-center my-2">
          <button
            className="disabled:opacity-40 py-2 px-4 rounded text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none"
            type="submit"
            disabled={!email || !password}
          >
            {isLogin ? 'Login' : 'Sign Up'}
          </button>
        </div>
      </form>
      <ArrowRightOnRectangleIcon
        className="my-2 h-6 w-6 cursor-pointer text-blue-500"
        onClick={() => setIsLogin(!isLogin)}
      />
    </div>
  )
}
