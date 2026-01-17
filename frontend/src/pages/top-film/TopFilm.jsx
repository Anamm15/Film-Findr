import { useFetchTopFilms, useFetchTrendingFilms } from "../../hooks/useFilm";
import ListFilm from "./components/ListFilm";


const TopFilmPage = () => {
   const { data: topFilms, loading: topFilmLoading } = useFetchTopFilms();
   const { data: trendingFilms, loading: trendingFilmLoading } = useFetchTrendingFilms();

   return (
      <>
         <div className="p-4 xl:max-w-[1280px] mx-auto mt-12">
            <h1 className="text-3xl md:text-4xl font-bold mb-4 text-text mt-10 font-geist">Top Film</h1>
            {
               topFilms && <ListFilm films={topFilms} loading={topFilmLoading} />
            }

            <h1 className="text-3xl md:text-4xl font-bold mb-4 text-text mt-10 font-geist">Trending Film</h1>
            {
               trendingFilms && <ListFilm films={trendingFilms} loading={trendingFilmLoading} />
            }
         </div>
      </>
   )
}

export default TopFilmPage;