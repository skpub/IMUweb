import type { MetaTagsProps } from 'svelte-meta-tags'
import { PUBLIC_BACKEND_ADDR } from '$env/static/public'

const imu = `${PUBLIC_BACKEND_ADDR}/IMU_minecraft.webp`

export const load = ({ url }) => {
  const baseMetaTags = Object.freeze({
    title: 'インモラル大学公式',
    titleTemplate: '%s | インモラル大学公式',
    description: 'インモラル大学の公式サイトです。',
    canonical: new URL(url.pathname, url.origin).href,
    twitter: {
      cardType: 'summary_large_image' as const,
      site: '@OMGR_dearinsu',
      creator: '@OMGR_dearinsu',
      title: 'インモラル大学公式',
      description: 'インモラル大学の公式サイトです',
      image: imu,
      imageAlt: 'インモラル大学のキャンパス'
    },
    openGraph: {
      type: 'website',
      url: new URL(url.pathname, url.origin).href,
      locale: 'ja_JP',
      title: 'インモラル大学公式',
      description: 'インモラル大学の公式サイトです',
      siteName: 'インモラル大学公式サイト',
      images: [
        {
          url: imu,
          alt: 'Og Image Alt',
          Width: 800,
          height: 600,
          secureUrl: imu,
          type: 'image/webp'
        }
      ]
    },
  }) satisfies MetaTagsProps

  return {
    baseMetaTags
  }
}
