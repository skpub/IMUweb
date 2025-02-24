import { PUBLIC_BACKEND_ADDR, PUBLIC_BACKEND_PORT } from "$env/static/public";
import { writable } from "svelte/store";

export let LoggedIn = writable<boolean | undefined>(undefined)

export async function reload() {
  try {
    const res = await fetch(PUBLIC_BACKEND_ADDR + ":" + PUBLIC_BACKEND_PORT + "/api/refresh", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      credentials: "include"
    })
    if (!res.ok) throw new Error("Failed to refresh session.")
    LoggedIn.set(true)
  } catch (e) {
    throw e
  }
}

export async function login(studentID: string, password: string) {
  try {
    const res = await fetch(PUBLIC_BACKEND_ADDR + ":" + PUBLIC_BACKEND_PORT + "/api/student/login", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        studentID: studentID,
        password: password
      })
    })
    if (!res.ok) throw new Error("Login failed.")
    LoggedIn.set(true)
  } catch (e) {
    throw e
  }
}

export async function logout() {
  try {
    const res = await fetch(PUBLIC_BACKEND_ADDR + ":" + PUBLIC_BACKEND_PORT + "/api/student/logout", {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
      credentials: "include"
    })
    if (!res.ok) throw new Error("Logout failed.")
    LoggedIn.set(false)
  } catch (e) {
    throw e
  }
}
