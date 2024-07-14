-- Create tables (order matters due to foreign keys)

CREATE TABLE artist (
    id TEXT PRIMARY KEY,
    artist_name TEXT NOT NULL,
    external_url TEXT,
    _checksum TEXT
);

CREATE TABLE album (
    id TEXT PRIMARY KEY,
    album_name TEXT NOT NULL,
    album_type TEXT,
    release_date TEXT, -- Assumes date format, adjust if different
    release_date_precision TEXT, -- "year", "month", etc.
    total_tracks INTEGER,
    external_url TEXT,
    _checksum TEXT
);

CREATE TABLE track (
    id TEXT PRIMARY KEY,
    track_name TEXT NOT NULL,
    disc_number INTEGER,
    duration_ms INTEGER,
    track_explicit BOOLEAN,
    is_local BOOLEAN,
    popularity INTEGER,
    preview_url TEXT,
    track_number INTEGER,
    external_url TEXT,
    _checksum TEXT,
    album_id TEXT REFERENCES album(id) ON DELETE CASCADE -- Link to album
);

CREATE TABLE image (
    image_url TEXT PRIMARY KEY, -- Assuming each image is unique
    height INTEGER,
    width INTEGER
);

-- Junction tables for many-to-many relationships

CREATE TABLE album_artists (
    album_id TEXT REFERENCES album(id) ON DELETE CASCADE,
    artist_id TEXT REFERENCES artist(id) ON DELETE CASCADE,
    PRIMARY KEY (album_id, artist_id) -- Each combo unique
);

CREATE TABLE track_artists (
    track_id TEXT REFERENCES track(id) ON DELETE CASCADE,
    artist_id TEXT REFERENCES artist(id) ON DELETE CASCADE,
    PRIMARY KEY (track_id, artist_id)
);

CREATE TABLE album_images (
    album_id TEXT REFERENCES album(id) ON DELETE CASCADE,
    image_url TEXT REFERENCES image(url) ON DELETE CASCADE,
    PRIMARY KEY (album_id, image_url)
);


CREATE TABLE activity (
    id SERIAL PRIMARY KEY,
    played_at TIMESTAMP NOT NULL,
    activity_type TEXT,
    external_url TEXT,
    track_id TEXT REFERENCES track(id) ON DELETE CASCADE
);
