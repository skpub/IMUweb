<script lang="ts">
  import { onMount } from "svelte";
  import { PUBLIC_BACKEND_ADDR, PUBLIC_BACKEND_PORT } from "$env/static/public";

  type Student = {
    studentID: string
    name: string
    bio: string
    img: BinaryType
  }

  let students: Student[] = $state([])

  onMount(() => {
    fetch(`${PUBLIC_BACKEND_ADDR}:${PUBLIC_BACKEND_PORT}/api/students/profile`)
      .then((res) => res.json())
      .then((data) => {
        data.forEach((student: Student) => {
          students.push(student)
        })
      });
      console.log(students)
  })
</script>

<div class="margin-case">
  <h1>学生一覧</h1>
  {#each students as student}
    <div>
      <img src={student.img} alt={student.name} />
      <h2>{student.name}</h2>
      <p>{student.bio}</p>
    </div>
  {/each}
</div>

<style>

</style>