import type { MetaTagsProps } from 'svelte-meta-tags'
import imu from '$lib/assets/IMU_minecraft.webp'

export const load = ({ url }) => {
  const baseMetaTags = Object.freeze({
    title: 'インモラル大学公式',
    titleTemplate: '%s | インモラル大学公式',
    description: 'インモラル大学の公式サイトです。',
    canonical: new URL(url.pathname, url.origin).href,
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
    }
  }) satisfies MetaTagsProps

  return {
    baseMetaTags
  }
}
