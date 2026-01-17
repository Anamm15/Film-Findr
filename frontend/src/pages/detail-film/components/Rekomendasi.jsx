import { useNavigate } from "react-router-dom";
import FilmCard from "../../../components/FilmCard";

const RekomendasiFilm = (props) => {
   const { films } = props;
   const navigate = useNavigate();

   const handleClickCard = (id) => {
      navigate(`/film/${id}`);
   };

   return (
      <>
         <h1 className="text-3xl font-bold mb-4 mt-8">Recommendation</h1>
         <div className="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-6">
            {films && films.slice(0, 4).map((film) => (
               <FilmCard key={film.id} movie={film} onClick={() => handleClickCard(film.id)} />
            ))}
         </div>
      </>
   )
}

export default RekomendasiFilm;