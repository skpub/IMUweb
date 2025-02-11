import type { PageServerLoad } from './$types';
import { BACKEND_ADDR, BACKEND_PORT } from '$env/static/private'

export const load: PageServerLoad = async ({ fetch, params }) => {
  const resList = await fetch(BACKEND_ADDR + ":" + BACKEND_PORT + "/api/article/list")
  const list = await resList.json()
  
  return { list }
}