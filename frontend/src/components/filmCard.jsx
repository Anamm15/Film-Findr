const FilmCard = ({ movie, onClick, children = null }) => {
   const {
      film_gambar,
      judul,
      genres,
      tanggal_rilis,
      rating,
      status,
      durasi,
   } = movie;

   return (
      <div
         className="max-w-sm bg-background rounded-2xl overflow-hidden shadow-lg hover:shadow-xl transition-shadow duration-300 cursor-pointer"
         onClick={onClick}
      >
         {/* Gambar */}
         <img
            className="w-full h-[260px] md:h-[320px] object-cover"
            src={film_gambar[0].url}
            alt={judul}
         />

         <div className="px-4 pt-2 pb-4">
            {/* Judul */}
            <h2 className="text-lg md:text-xl font-semibold mb-1 truncate text-text">
               {judul}
            </h2>

            {/* Genre hanya tampil di tablet ke atas */}
            <div className="hidden md:flex flex-wrap gap-2 text-sm text-white mb-2">
               {genres &&
                  genres.map((genre) => (
                     <span
                        key={genre.id}
                        className="bg-tertiary px-2 py-0.5 rounded-full"
                     >
                        {genre.nama}
                     </span>
                  ))}
            </div>

            {/* Detail info (hanya tablet ke atas) */}
            <div className="hidden md:block text-sm text-gray-600 space-y-1">
               <p>
                  <strong>Release Date:</strong> {tanggal_rilis}
               </p>
               <p>
                  <strong>Duration:</strong> {durasi} minutes
               </p>
               <p>
                  <strong>Status:</strong>{" "}
                  <span className="capitalize">{status}</span>
               </p>
               <p>
                  <strong>Rating:</strong> ⭐ {rating}/10
               </p>
            </div>

            {/* Versi smartphone: tampilkan ringkas */}
            <div className="flex justify-between items-center text-sm text-gray-600 md:hidden mt-1">
               <span className="capitalize">{status}</span>
               <span>⭐ {rating}</span>
            </div>

            {children}
         </div>
      </div>
   );
};

export default FilmCard;
