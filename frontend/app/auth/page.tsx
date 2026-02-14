"use client"
import { useAuth } from '@/hooks/useAuth'
import { FormEvent, useState } from 'react'
import { AuthForm } from '@/components/AuthForm'

export default function Auth() {
  const { loginMutation, registerMutation } = useAuth()
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const [isLogin, setIsLogin] = useState(true)
  const isLoading = loginMutation.isPending || registerMutation.isPending

  const handleSubmit = async (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    isLogin 
      ? loginMutation.mutate({email, password}) 
      : registerMutation.mutate({email, password})
  }

  return (
    <div className="flex justify-center items-center flex-col min-h-screen text-gray-600 font-mono">
      <div className="flex items-center">
        <span className="text-center text-3xl font-extrabold">
          Todo
        </span>
      </div>
      <h2 className="my-6">{isLogin ? 'ログイン' : '新規登録'}</h2>
      <AuthForm
        handleSubmit={handleSubmit}
        email={email}
        setEmail={setEmail}
        password={password}
        setPassword={setPassword}
        isLogin={isLogin}
        isSending={isLoading}
      />
      <p className="text-sm text-blue-500 cursor-pointer font-medium mt-4 underline hover:text-blue-600" onClick={() => setIsLogin(!isLogin)}>{isLogin ? 'アカウント登録はこちら' : 'ログインはこちら'}</p>
    </div>
  )
}
