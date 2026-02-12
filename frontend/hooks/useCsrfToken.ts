import { useEffect } from 'react'
import { api, csrf } from '@/lib/axios'

export const useCsrfToken = () => {
  useEffect(() => {
    const fetchToken = async () => {
      // Avoid fetching if already set (though checking header directly is tricky if axios instance shared)
      api.defaults.headers.common['X-CSRF-Token'] || await csrf()
    }
    fetchToken()
  }, [])
}
