import { useState, useEffect } from "react";
import { createFilm } from "../../../service/film";
import { getAllGenre } from "../../../service/genre";
import Input from "../../../components/Input";
import Button from "../../../components/Button";
import TextArea from "../../../components/Textarea";

const AddFilmPage = () => {
   const [judul, setJudul] = useState("");
   const [status, setStatus] = useState("");
   const [sinopsis, setSinopsis] = useState("");
   const [sutradara, setSutradara] = useState("");
   const [tanggalRilis, setTanggalRilis] = useState("");
   const [durasi, setDurasi] = useState(null);
   const [totalEpisode, setTotalEpisode] = useState(null);
   const [imageFiles, setImageFiles] = useState([]);
   const [genres, setGenres] = useState([]);
   const [selectedGenres, setSelectedGenres] = useState([]);
   const [message, setMessage] = useState("");
   const [colorMessage, setColorMessage] = useState("");
   const [isGenresLoaded, setIsGenresLoaded] = useState(false);


   useEffect(() => {
      const fetchGenres = async () => {
         try {
            const response = await getAllGenre();
            setGenres(response.data.data);
         } catch (error) {
            console.error("Error fetching genres:", error.data.message);
         }
      };

      fetchGenres();
      if (!isGenresLoaded) {
         setIsGenresLoaded(true);
      }
   }, [isGenresLoaded]);

   const handleSubmit = async (e) => {
      e.preventDefault();
      try {
         const formData = new FormData();
         formData.append("judul", judul);
         formData.append("status", status);
         formData.append("sinopsis", sinopsis);
         formData.append("sutradara", sutradara);
         formData.append("tanggal_rilis", tanggalRilis);
         formData.append("durasi", durasi);
         formData.append("total_episode", totalEpisode);
         selectedGenres.forEach((g) => formData.append("genres", g));

         for (let i = 0; i < imageFiles.length; i++) {
            formData.append("images", imageFiles[i]);
         }

         const response = await createFilm(formData);
         if (response.status === 201) {
            setMessage(response.data.message);
            setColorMessage("text-green-600");
         }
      } catch (error) {
         console.log(error);
         setMessage(error.data.message);
         setColorMessage("text-red-600");
      }
   };

   return (
      <>
         <h1 className="text-4xl font-bold mb-5">Add Film</h1>
         <form
            onSubmit={handleSubmit}
            className="w-full relative mb-10">
            <div className="w-full grid lg:grid-cols-2 gap-4">
               <Input
                  type="text"
                  placeholder="Masukkan Judul"
                  label="Judul"
                  value={judul}
                  onChange={(e) => setJudul(e.target.value)}
               />
               <Input
                  type="text"
                  label="Status"
                  placeholder="Pilih Status"
                  value={status}
                  onChange={(e) => setStatus(e.target.value)}
               />
               <Input
                  type="text"
                  label="Sutradara"
                  placeholder="Masukkan Sutradara"
                  value={sutradara}
                  onChange={(e) => setSutradara(e.target.value)}
               />
               <Input
                  type="date"
                  label="Tanggal Rilis"
                  placeholder="Masukkan tanggal rilis"
                  value={tanggalRilis}
                  onChange={(e) => setTanggalRilis(e.target.value)}
               />
               <Input
                  type="number"
                  label="Total Episode"
                  placeholder="Masukkan Total Episode"
                  value={totalEpisode}
                  onChange={(e) => setTotalEpisode(e.target.value)}
               />
               <Input
                  type="number"
                  label="Durasi"
                  placeholder="Masukkan Durasi"
                  value={durasi}
                  onChange={(e) => setDurasi(e.target.value)}
               />
               <Input
                  type="file"
                  label="Image"
                  multiple={true}
                  onChange={(e) => setImageFiles(e.target.files)}
               />
            </div>
            <TextArea
               label="Sinopsis"
               placeholder="Masukkan Sinopsis"
               value={sinopsis}
               onChange={(e) => setSinopsis(e.target.value)}
            />
            <div className="mt-4">
               <label className="block text-gray-600 text-sm mb-2">Pilih Genre</label>
               <div className="flex flex-wrap gap-2">
                  {genres.map((genre) => (
                     <button
                        key={genre.id}
                        type="button"
                        className={`px-3 py-1 text-sm rounded-full border transition
                        ${selectedGenres.includes(genre.id)
                              ? "bg-primary text-white border-primary"
                              : "bg-background text-gray-700 border-gray-300 hover:bg-gray-100"}`}
                        onClick={() => {
                           setSelectedGenres((prev) =>
                              prev.includes(genre.id)
                                 ? prev.filter((id) => id !== genre.id)
                                 : [...prev, genre.id]
                           );
                        }}
                     >
                        {genre.nama}
                     </button>
                  ))}
               </div>
            </div>
            <Button
               type="submit"
               className="mt-4 w-40 py-2 rounded-lg absolute right-0"
            >
               Submit
            </Button>
            {message && <p className={`mt-16 absolute right-0 text-sm ${colorMessage}`}>{message}</p>}
         </form>
      </>
   )
}

export default AddFilmPage;