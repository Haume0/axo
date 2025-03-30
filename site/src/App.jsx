import { useEffect, useState } from "react";

function App() {
  const [exampleUrl, setExampleUrl] = useState("/api/hello");
  const [mloading, _mloading] = useState(false);
  const [notes, setNotes] = useState([]);
  useEffect(() => {
    const fetchNotes = async () => {
      const res = await fetch("/api/notes");
      if (res.ok) {
        const data = await res.json();
        setNotes(data);
      } else {
        console.error("Failed to fetch notes");
      }
    };
    fetchNotes();
  }, []);
  return (
    <main className="min-w-[27rem] max-w-[34rem] shrink-0 mx-auto relative flex flex-col items-center gap-4 justify-center p-8 text-center">
      <div>
        <a href="https://haume.me/axo" target="_blank">
          <img
            src="/api/static/axo.webp"
            alt="Vite logo"
            className="h-32 transition-all hover:drop-shadow-[0_0_2em_#646cffaa]"
          />
        </a>
      </div>
      <h1 className="text-3xl font-bold">🪸 Welcome to Axo 🌊</h1>
      <p className="">
        AxoScaffold is a Restful API scaffold for Go, built on top of stdlib and
        gorm. <br /> It is designed to be simple, fast, and easy to use. For
        more information, please click the logo above.
      </p>
      <p className="dark:text-white/40 text-black/20 text-sm">
        License: MIT <br />
        Copyright (c) 2025 Haume
      </p>
      <div className="flex justify-center items-center  gap-3">
        <button
          onClick={() => {
            setExampleUrl("/api/hello");
          }}>
          Hello Example
        </button>
        <button
          onClick={() => {
            setExampleUrl("/api/error");
          }}>
          Error Example
        </button>
        <button
          onClick={() => {
            setExampleUrl(
              "/api/image?src=axo.webp&format=jpeg&quality=80&width=200&height=200"
            );
          }}>
          Image Optimization
        </button>
      </div>
      <div className="inline-flex justify-center items-center gap-3 w-full">
        <input
          type="text"
          name="url"
          value={exampleUrl}
          readOnly
          className="w-full h-12 rounded-lg px-4 text-sm font-mono bg-black/10 dark:bg-black/20 outline-none border-none"
        />
        <a href={exampleUrl} target="_blank">
          <button>Open</button>
        </a>
      </div>
      <iframe
        src={exampleUrl}
        className="rounded-2xl w-full h-[27rem] p-4 bg-black/10 dark:bg-black/20"
      />
      <hr />
      <h4 className="text-2xl font-bold">Test Mail!</h4>
      <form
        onSubmit={async (e) => {
          e.preventDefault();
          _mloading(true);
          let formdata = new FormData(e.currentTarget);
          let mail = formdata.get("email");
          const res = await fetch(`/api/testmail?mail=${mail}`);
          if (res.ok) {
            const a = await res.json();
            alert(a.message);
          } else {
            const a = await res.json();
            alert(a.error);
          }
          _mloading(false);
        }}
        className={`flex gap-2 w-full ease-in-out duration-700 ${
          mloading && "scale-95 blur-[2px] opacity-60"
        }`}>
        <input
          className="w-full h-12 rounded-lg px-4 text-sm bg-black/10 dark:bg-black/20 outline-none border-none"
          type="email"
          name="email"
          placeholder="Enter a mail adress!"
          id="email"
        />
        <button>Send</button>
      </form>
      <hr />
      <h4 className="text-2xl font-bold">Notes App!</h4>
      <form
        onSubmit={async (e) => {
          e.preventDefault();
          let formdata = new FormData(e.currentTarget);
          let note = formdata.get("Note");
          const res = await fetch(`/api/notes?note=${note}`, {
            method: "POST",
          });
          if (res.ok) {
            const data = await res.json();
            setNotes((prev) => [...prev, data]);
          }
          // Reset the Note input field after successful submission
          e.target.elements.Note.value = "";
        }}
        className="flex gap-2 w-full">
        <input
          type="text"
          name="Note"
          placeholder="Enter a note!"
          className="w-full h-12 rounded-lg px-4 text-sm font-mono bg-black/10 dark:bg-black/20 outline-none border-none"
        />
        <button>Add</button>
      </form>
      <ul className="w-full min-h-[24rem] rounded-2xl gap-2 flex flex-col p-2 text-sm bg-black/10 dark:bg-black/20 outline-none border-none">
        <p className="font-semibold text-xs text-center text-white/40">
          Note List
        </p>
        {notes.map((note, index) => (
          <li
            key={index}
            className="w-full p-1 rounded-xl flex items-center justify-between text-sm bg-black/10 dark:bg-black/20 outline-none border-none">
            <p className="px-4">{note.title}</p>
            <button
              onClick={() => {
                fetch(`/api/notes?id=${note.id}`, {
                  method: "DELETE",
                }).then((res) => {
                  if (res.ok) {
                    setNotes(notes.filter((n) => n.id !== note.id));
                  } else {
                    // alert("Failed to delete note");
                  }
                });
              }}>
              Delete
            </button>
          </li>
        ))}
      </ul>
    </main>
  );
}

export default App;
