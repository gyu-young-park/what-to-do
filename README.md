# What-to-do
업무가 너무 많을 때 한 가지씩 업무를 미루다보면 잊어버리는 경우가 생긴다. 더 이상 내 기억력을 믿지말고 기록해두며 살도록 하자.

그럼 스티커 메모나 노션같은 프로그램에 적어두면 되지 않는가?? 마우스도 쓰기 싫고, terminal에서 다 해결되면 좋겠다. 그럼 CLI로 todo list를 만들기로 하자.

# 기본 명령어
1. `-add {task}`: `task`를 `todo.json`에 추가합니다. ex) `-add study golang grammar`라고 추가하면 `study golang grammar`가 `task`가 추가됩니다.
2. `-list`: `task` list를 보여줍니다. ex) 
```
╔═══╤══════════════════╤═══════╤═════════════════════╤═════════════════════╗
║ # │       Task       │ Done? │           CreatedAt │         CompletedAt ║
╟━━━┼━━━━━━━━━━━━━━━━━━┼━━━━━━━┼━━━━━━━━━━━━━━━━━━━━━┼━━━━━━━━━━━━━━━━━━━━━╢
║ 1 │ Sample test      │ no    │ 25 Jun 22 19:20 KST │ 01 Jan 01 00:00 UTC ║
║ 2 │ ✅ Sample test   │ yes   │ 25 Jun 22 19:20 KST │ 25 Jun 22 20:06 KST ║
║ 3 │ Sample test      │ no    │ 25 Jun 22 19:20 KST │ 01 Jan 01 00:00 UTC ║
║ 4 │ hello world      │ no    │ 25 Jun 22 19:39 KST │ 01 Jan 01 00:00 UTC ║
║ 5 │ list up my files │ no    │ 25 Jun 22 19:39 KST │ 01 Jan 01 00:00 UTC ║
║ 6 │ hello worlds     │ no    │ 26 Jun 22 01:51 KST │ 01 Jan 01 00:00 UTC ║
╟━━━┼━━━━━━━━━━━━━━━━━━┼━━━━━━━┼━━━━━━━━━━━━━━━━━━━━━┼━━━━━━━━━━━━━━━━━━━━━╢
║                       you have [5] pending todos!                        ║
╚═══╧══════════════════╧═══════╧═════════════════════╧═════════════════════╝
```
완료한 일은 <span style="color:green;">초록색</span>과 체크 표시로 되고, 아직 완료되지 못한 일은 <span style="color:blue;">파란색</span>으로 표시됩니다.

3. `-complete={number}`: `number`번째 `task`를 완료 표시로 둡니다. ex) 위 리스트에서 `list up my files`를 완료 표시로 두고 싶다면, `-complete=5`로 두시면 됩니다.
4. `-delete={number}`: `number`번째 `task`를 리스트에서 삭제 합니다. ex) 위 리스트에서 `list up my files`를 삭제하고 싶다면, `-delete=5`로 두시면 됩니다.

# pipe 기능
`add` 명령에 대해서 linux pipe를 지원합니다. pipe를 사용하는 경우는 `-add`뒤에 `{task}`를 입력하지 않아야 합니다. `-add`뒤에 `{task}`를 입력하면 pipe로 넘어간 bytes는 무시하고 `{task}`가 우선됩니다.
```
echo "new Task" | go run ./cmd/todo/ -add
```
위와 같이 pipe를 사용하여 기존에 할 일을 메모해둔 파일이 있다면 이를 사용하여 할 일을 추가하시면 됩니다.
