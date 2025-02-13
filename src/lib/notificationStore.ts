import { writable } from 'svelte/store'
import Notification from './Notification.svelte'

export const notifications = writable<OneNotification[]>([])

export type NotificationType = "info" | "error" | "warning"
export type OneNotification = {
  msg: string
  type: NotificationType
}

export function notify(msg: string, type: NotificationType ,duration = 2000) {
  const newNot: OneNotification = { msg, type }
  notifications.update(v => [...v, newNot])
  setTimeout(() => notifications.update(v => {
    let temp = v.slice()
    temp.reverse()
    temp.pop()
    temp.reverse()
    return temp
  }), duration)
}
