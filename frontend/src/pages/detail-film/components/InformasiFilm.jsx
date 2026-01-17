const InformasiFilm = (props) => {
    const { film } = props;
    return (
        <div className="bg-white mb-6 space-y-3 md:text-lg text-secondary">
            <h1 className="text-[26px] md:text-3xl font-bold mb-4 text-text">{film.judul}</h1>
            <div className="flex flex-wrap gap-2 mb-2">
                {film.genres?.map((genre) => (
                    <span
                        key={genre.id}
                        className="bg-tertiary text-white px-3 py-1 rounded-full"
                    >
                        {genre.nama}
                    </span>
                ))}
            </div>
            <p><strong>Release Date:</strong> {film.tanggal_rilis}</p>
            <p><strong>Duration:</strong> {film.durasi} menit</p>
            <p><strong>Status:</strong> <span className="capitalize">{film.status}</span></p>
            <p><strong>Rating:</strong> ⭐ {film.rating}/10</p>
            <p><strong>Director:</strong> {film.sutradara}</p>
            <p><strong>Total Episode:</strong> {film.total_episode}</p>
        </div>
    )
}

export default InformasiFilm