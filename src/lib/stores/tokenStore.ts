import { PUBLIC_BACKEND_ADDR, PUBLIC_BACKEND_PORT } from "$env/static/public";
import { writable } from "svelte/store";

type UserInfo = {
  studentID: string
  token: string
}

export let LoggedIn = writable<UserInfo | undefined>()
LoggedIn.subscribe(value => {
  if (value !== undefined)
  setCookie(value.studentID, value.token)
})

export const loadCookie = () => {
  const cookie = getCookie()
  if (cookie !== null) {
    LoggedIn.set(cookie)
  }
}

export const unloadCookie = () => {
  document.cookie = "studentID=; max-age=0"
  document.cookie = "token=; max-age=0"
}

type Cookie = {
  studentID: string
  token: string
}

function setCookie(studentID: string, token: string) {
  const exp = Date.now() + 1000 * 60 * 5
  document.cookie = `studentID=${studentID}; expires=${exp}`
  document.cookie = `token=${token}; expires=${exp}`
}

function getCookie(): Cookie | null {
  const cookie = document.cookie.split(';').map(c => c.trim())
  const studentID = cookie.find(c => c.startsWith('studentID'))?.split('=')[1]
  const token = cookie.find(c => c.startsWith('token'))?.split('=')[1]
  if (studentID === undefined || token === undefined) {
    return null
  } else {
    return { studentID: studentID, token: token }
  }
}

//
// static
//
export class Login {
  private static timer: ReturnType<typeof setInterval> | undefined = undefined

  private static async refresh() {
    let newToken: string | undefined
    let currentState: UserInfo | undefined

    LoggedIn.update(v => {
      currentState = v
      return v
    })

    if (!currentState) {
      throw new Error("current state is undefined.")
    }

    try {
      const res = await fetch(PUBLIC_BACKEND_ADDR + ":" + PUBLIC_BACKEND_PORT + "/api/refresh", {
        method: "POST",
        headers: {
          "Authorization": currentState.token,
          "Content-Type": "application/json"
        },
        body: JSON.stringify({ "token": currentState.token })
      })

      if (!res.ok) {
        throw new Error("failed to refresh")
      }

      const data = await res.json()
      newToken = data["token"]

    } catch (e) {
      // ネットワークがおかしいとき(など)はここに来る。
      throw e
    }
    LoggedIn.set({ studentID: currentState.studentID, token: newToken! })
  }

  private static async login_(studentname: string, password: string) {
    try {
      const res = await fetch(PUBLIC_BACKEND_ADDR + ":" + PUBLIC_BACKEND_PORT + "/api/student/login", {
        method: "POST",
        headers: {
          "Content-Type": "application/json"
        },
        body: JSON.stringify({ "studentID": studentname, "password": password })
      })
      if (!res.ok) {
        throw new Error("Failed to login")
      }
      const data = await res.json()
      const userinfo: UserInfo = {
        studentID: studentname,
        token: data
      }
      LoggedIn.set(userinfo)
    } catch (e) {
      // ネットワークがおかしいとき(など)はここに来る。
      throw e
    } 
  }
  // こっから下は公開のやつら

  public static async login(studentName: string, password: string) {
    // n>1回目の呼び出しだったときにtimerが動いていたら重複するので殺しておく。
    Login.destroy()
    try {
      await Login.login_(studentName, password)
      Login.timer = setInterval(() => {
        Login.refresh()
      }, (1000 * 60 * 4 + 1000 * 50)) // 5分でトークンが切れるので、4分50秒でリフレッシュ。
      // }, (10 * 60 * 4 + 10 * 50)) // 5 / 100 分でトークンが切れるので、4分50秒でリフレッシュ。
    } catch (e) {
      throw e;
    }
  }

  public static async restart_refresh() {
    const cookie = getCookie()
    if (cookie != null) {
      Login.destroy()
      this.refresh()
      Login.timer = setInterval(() => {
        Login.refresh()
      }, (1000 * 60 * 4 + 1000 * 50)) // 5分でトークンが切れるので、4分50秒でリフレッシュ。
      // }, (10 * 60 * 4 + 10 * 50)) // 5 / 100 分でトークンが切れるので、4分50秒でリフレッシュ。
    }
  }

  public static destroy() {
    clearInterval(Login.timer)
    Login.timer = undefined
  }
}
