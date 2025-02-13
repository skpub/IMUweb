<script lang="ts">
  import { LoggedIn, Login } from "$lib/stores/tokenStore";
  import { notify } from "$lib/notificationStore";
  import { goto } from "$app/navigation";
  import { page } from "$app/state";

  let redirect: string | null = page.url.searchParams.get('redirect')
  let student_id = '';
  let password = '';
  async function login() {
    try {
      await Login.login(student_id, password)
      notify("ログインに成功しました", "info")
      goto(redirect ?? '/')
    } catch (e) {
      notify("ログインに失敗しました", "error")
    }
  }
</script>

<div id="login_dialog">
  <h1>認証</h1>
  <form onsubmit={login}>
    <input bind:value={student_id} type="text" name="username" placeholder="学生・教員ID" required>
    <input bind:value={password} type="password" name="password" placeholder="パスワード" required>
    <button type="submit">ログイン</button>
  </form>
</div>

<style>
  #login_dialog {
    margin: 0 auto;
    margin-top: 30px;
    width: 200px;
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