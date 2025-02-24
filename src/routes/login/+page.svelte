<script lang="ts">
  import { login, LoggedIn } from "$lib/stores/session";
  import { notify } from "$lib/notificationStore";
  import { goto } from "$app/navigation";
  import { page } from "$app/state";

  let redirect: string | null = page.url.searchParams.get('redirect')
  let student_id = '';
  let password = '';
  async function login_() {
    try {
      await login(student_id, password)
      notify("ログインに成功しました", "info")
      LoggedIn.set(true)
      goto(redirect ?? '/')
    } catch (e) {
      notify("ログインに失敗しました", "error")
    }
  }
  function uraguchi() {
    goto('/uraguchi')
  }
</script>

<div id="login_dialog">
  <h1>認証</h1>
  <form onsubmit={login_}>
    <input bind:value={student_id} type="text" name="username" placeholder="学生・教員ID" required>
    <input bind:value={password} type="password" name="password" placeholder="パスワード" required>
    <button type="submit">ログイン</button>
  </form>
  <button id="uraguchi" onclick={uraguchi}>裏口入学 (違法登録)</button>
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
    #uraguchi {
      margin-top: 40px;
    }
  }
</style>
