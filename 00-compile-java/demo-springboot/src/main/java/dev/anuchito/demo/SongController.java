package dev.anuchito.demo;

import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;

@RestController
public class SongController {

	@GetMapping("/songs")
	public Song getLoan() {
		Song song = new Song();
		song.setSongId("L001");
		song.setSongName("Laumcing");
		return song;
	}
}
