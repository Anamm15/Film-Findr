import { useState } from "react";
import { searchFilm } from "../../service/film";
import ListFilm from "./components/ListFilm";
import Button from "../../components/Button";
import { useFetchFilms } from "../../hooks/useFilm";
import { useEffect } from "react";
import { INIT_PAGE_NUMBER } from "../../utils/constant";

const LandingPage = () => {
   const [page, setPage] = useState(INIT_PAGE_NUMBER);
   const { data, loading } = useFetchFilms(page);
   const [films, setFilms] = useState(data);
   const [searchQuery, setSearchQuery] = useState("");

   useEffect(() => {
      if (!data) return;
      setFilms(data)
   }, [data]);

   const handleSearch = async () => {
      try {
         const response = await searchFilm(searchQuery);
         setFilms(response.data.data);
      } catch (error) {
         console.error("Error searching films:", error);
         setFilms([]);
      }
   };

   return (
      <>
         <div className="p-4 xl:max-w-[1280px] mx-auto mt-28">
            <div className="mb-6 relative max-w-2xl mx-auto flex gap-2">
               <div className="relative w-full">
                  <input
                     type="text"
                     value={searchQuery}
                     onChange={(e) => setSearchQuery(e.target.value)}
                     onKeyDown={(e) => {
                        if (e.key === "Enter") handleSearch();
                     }}
                     placeholder="🔍 Search film you interested..."
                     className="w-full py-3 pl-5 pr-12 rounded-xl border border-gray-300 shadow-md focus:outline-none focus:ring-2 focus:ring-primary focus:border-transparent transition-all duration-300 bg-white text-gray-700 md:text-lg"
                  />
                  {searchQuery && (
                     <button
                        onClick={() => setSearchQuery("")}
                        className="absolute right-4 top-1/2 -translate-y-1/2 text-gray-400 hover:text-gray-600 text-2xl"
                     >
                        ×
                     </button>
                  )}
               </div>
               <Button
                  onClick={handleSearch}
                  className="px-2 md:px-6 py-3 flex items-center text-xl rounded-xl bg-gradient-primary text-white font-semibold hover:bg-opacity-90 transition-all shadow gap-1"
               >
                  <span>🔍</span>
                  <span className="hidden md:inline">Cari</span>
               </Button>
            </div>

            <h1 className="text-3xl md:text-4xl font-bold mb-4 text-text font-geist">Last Film</h1>
            {
               films && <ListFilm films={films} page={page} setPage={setPage} loading={loading} />
            }
            {
               films?.length === 0 && (
                  <h1 className="text-lg mb-4 w-full text-center font-geist">Film not found</h1>

               )
            }
         </div>
      </>
   );
}

export default LandingPage;