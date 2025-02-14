import type { PageServerLoad } from './$types';
import { PRIVATE_BACKEND_ADDR, PRIVATE_BACKEND_PORT } from '$env/static/private'

export const load: PageServerLoad = async ({ fetch, params }) => {
  const resList = await fetch(PRIVATE_BACKEND_ADDR + ":" + PRIVATE_BACKEND_PORT + "/api/article/list")
  const list = await resList.json()
  
  return { list }
}
