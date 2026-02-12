"use client"
import { useAuth } from '@/hooks/useAuth'
import { FormEvent, useState } from 'react'
import { CheckBadgeIcon, ArrowRightOnRectangleIcon, ArrowPathIcon } from '@heroicons/react/24/solid'
import { AuthForm } from '@/components/AuthForm'

export default function Auth() {
  const { loginMutation, registerMutation } = useAuth()
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const [isLogin, setIsLogin] = useState(true)
  const [isSending, setIsSending] = useState(false)

  const handleSubmit = async (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    setIsSending(true)
    if (isLogin) {
      loginMutation.mutate({
        email,
        password,
      })
    } else {
      registerMutation.mutate({
        email,
        password,
      })
    }
    setIsSending(false)
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
      <AuthForm
        handleSubmit={handleSubmit}
        email={email}
        setEmail={setEmail}
        password={password}
        setPassword={setPassword}
        isLogin={isLogin}
        isSending={isSending}
      />
      <p className="text-sm text-blue-500 cursor-pointer" onClick={() => setIsLogin(!isLogin)}>{isLogin ? 'Don\'t have an account?' : 'Already have an account?'}</p>
    </div>
  )
}
