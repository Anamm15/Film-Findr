const Sinopsis = (props) => {
    const { sinopsis } = props;
    return (
        <div className="bg-gray-100 rounded-xl p-6 mb-8 max-h-96 overflow-auto">
            <h2 className="text-2xl md:text-3xl font-semibold mb-2">Synopsis</h2>
            <p className="text-text md:text-xl whitespace-pre-wrap text-justify">{sinopsis}</p>
        </div>
    )
}

export default Sinopsis;