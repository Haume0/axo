import { createFileRoute, Link } from "@tanstack/react-router";
import { useState, useEffect } from "react";

export const Route = createFileRoute("/database")({
  component: RouteComponent,
});

interface Note {
  id: number;
  title: string;
}

function RouteComponent() {
  const [mloading, _mloading] = useState(false);
  const [notes, setNotes] = useState<Note[]>([]);
  useEffect(() => {
    const fetchNotes = async () => {
      const res = await fetch("/api/notes");
      if (res.ok) {
        const data = await res.json();
        setNotes(data);
      } else {
        // Handle error
        setNotes([]);
        console.error("Failed to fetch notes");
      }
    };
    fetchNotes();
  }, []);
  return (
    <>
      <h1 className="text-3xl font-bold">🪸 Axo, Gorm & PostgreSQL 🌊</h1>
      <p className="">
        This is a simple example of how to use Axo with Gorm and PostgreSQL.
        <br />
        It demonstrates how to create, read, update, and delete notes in a
        database.
      </p>
      <a href="https://gorm.io/docs/" className="text-accent hover:underline">
        Gorm Document →
      </a>
      <Link to="/">
        <button>← Back to Home</button>
      </Link>
      <h4 className="text-2xl font-bold">Notes App!</h4>
      <form
        onSubmit={async (e) => {
          e.preventDefault();
          _mloading(true);
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
          e.currentTarget.reset();
          _mloading(false);
        }}
        className={`flex gap-2 w-full ease-in-out duration-700 ${
          mloading && "scale-95 blur-[2px] opacity-60"
        }`}>
        <input type="text" name="Note" placeholder="Enter a note!" />
        <button>Add</button>
      </form>
      <ul className="w-full min-h-[24rem] rounded-xl gap-2 flex flex-col p-2 text-sm bg-comet-500 outline-none border-none">
        <p className="font-semibold text-xs text-center text-white/40">
          Note List
        </p>
        {notes &&
          notes.map((note, index) => (
            <li
              key={index}
              className="w-full p-1 rounded-xl flex items-center justify-between text-sm bg-comet-600 outline-none border-none">
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
    </>
  );
}
