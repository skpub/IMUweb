<script lang="ts">
  // import { LoggedIn, Login } from "$lib/stores/tokenStore";
  import { notify } from "$lib/notificationStore";
  import { goto } from "$app/navigation";
  import { page } from "$app/state";
  import { PUBLIC_BACKEND_ADDR, PUBLIC_BACKEND_PORT } from "$env/static/public";

  let redirect: string | null = page.url.searchParams.get('redirect')
  let student_id = '';
  let password = '';
  let name = '';
  let email = '';
  async function login() {
    try {
      const res = await fetch(PUBLIC_BACKEND_ADDR + ':' + PUBLIC_BACKEND_PORT + '/api/student/signup', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({
          "studentID": student_id,
          "password": password,
          "name": name,
          "email": email
        })
      })
      if (res.status !== 200) {
        notify("裏口入学に失敗しました", "error")
        return
      }
      notify("裏口入学に成功しました", "info")
      goto('/login')
    } catch (e) {
      notify("裏口入学に失敗しました(多分サーバが死んでる)", "error")
    }
  }
  function uraguchi() {
    goto('/uraguchi')
  }
</script>

<div id="login_dialog">
  <h1>裏口入学</h1>
  <form onsubmit={login}>
    <input bind:value={student_id} type="text" name="username" placeholder="学生・教員ID" required>
    <input bind:value={name} type="text" name="name" placeholder="名前" required>
    <input bind:value={email} type="text" name="email" placeholder="E-Mail" required>
    <input bind:value={password} type="password" name="password" placeholder="パスワード" required>
    <button type="submit">違法登録</button>
  </form>
</div>

<style>
  #login_dialog {
    margin: 0 auto;
    margin-top: 30px;
    width: 200px;
    display: flex;
    flex-flow: column;
    h1 {
      text-align: center;
    }
    form {
      display: flex;
      flex-direction: column;
      input {
        padding: 2px 2px 2px 10px;
        border: none;
        border-radius: 32px;
        margin: 4px 0;
        font-size: 18px;
        font-family: monospace;
        margin-top: 8px;
      }
      button {
        border: none;
        font-size: 24px;
        margin: 4px 0;
        border-radius: 24px;
        color: var(--white);
        background-color: var(--immoral-shadow);
      }
    }
  }
</style>