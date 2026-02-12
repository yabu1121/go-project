import { ArrowPathIcon } from "@heroicons/react/24/solid"

type AuthFormProps = {
  handleSubmit: (e: React.FormEvent) => void
  email: string
  setEmail: (e: string) => void
  password: string
  setPassword: (e: string) => void
  isLogin: boolean
  isSending: boolean
}

export const AuthForm = ({
  handleSubmit,
  email,
  setEmail,
  password,
  setPassword,
  isLogin,
  isSending,
}: AuthFormProps) => {
  return (
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

          {isLogin ? <div className="flex items-center">
            {isSending ? <ArrowPathIcon className="animate-spin h-5 w-5 mr-2" /> : <p>Login</p>}
          </div> : <div className="flex items-center">
            {isSending ? <ArrowPathIcon className="animate-spin h-5 w-5 mr-2" /> : <p>Sign Up</p>}
          </div>}
        </button>
      </div>
    </form>
  )
}