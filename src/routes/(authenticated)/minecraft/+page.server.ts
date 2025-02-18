import type { PageServerLoad } from './$types';
import { env } from '$env/dynamic/private'

export const load: PageServerLoad = async () => {
  return {
    JEAddr: env.PRIVATE_MC_JE_ADDR,
    BEAddr: env.PRIVATE_MC_BE_ADDR,
    version: env.PRIVATE_MC_VERSION,
  }
}
