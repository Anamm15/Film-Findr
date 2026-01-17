import { Link } from "react-router-dom";

const WatchlistLayout = (props) => {
    const { watchlist, children } = props;
    return (
        <div
            key={watchlist.id}
            className="flex gap-6"
        >
            <div className="flex-shrink-0 w-32 h-48 overflow-hidden rounded-xl shadow">
                <Link to={`/film/${watchlist.film.id}`}>
                    <img
                        src={watchlist.film.film_gambar[0].url}
                        alt={watchlist.film.judul}
                        className="w-full h-full object-cover"
                    />
                </Link>
            </div>

            <div className="flex flex-col justify-between w-full">
                {children}
            </div>
        </div>
    )
}

export default WatchlistLayout;