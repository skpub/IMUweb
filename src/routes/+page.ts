import type { PageLoad } from './$types';
import { PUBLIC_BACKEND_ADDR, PUBLIC_BACKEND_PORT } from '$env/static/public'

export const load: PageLoad = async ({ fetch, params }) => {
  const resList = await fetch(PUBLIC_BACKEND_ADDR + ":" + PUBLIC_BACKEND_PORT + "/api/article/list")
  const list = await resList.json()
  
  return { list }
}
