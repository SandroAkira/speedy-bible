var all_words;
var paused_queue = [];

// inspiration from [OpenSpritz](https://github.com/Miserlou/OpenSpritz)
function rsvpify(input_el, wpm){
  var input = $(input_el).text();
  var words_per_minute = wpm;
  var ms_per_word = 60000/wpm;

  all_words = input.split(/\s+/);

  var word = '';
  var result = '';
  var temp_words = all_words.slice(0); // copy Array
  var t = 0;

  for (var i=0; i<all_words.length; i++){

    // Double up on long words and words with commas.
    if((all_words[i].indexOf(',') != -1 || all_words[i].indexOf(':') != -1 || all_words[i].indexOf('-') != -1 || all_words[i].indexOf('(') != -1|| all_words[i].length > 8) && all_words[i].indexOf('.') == -1){
      temp_words.splice(t+1, 0, all_words[i]);
      temp_words.splice(t+1, 0, all_words[i]);
      t++;
      t++;
    }

    // Add an additional space after punctuation.
    if(all_words[i].indexOf('.') != -1 || all_words[i].indexOf('!') != -1 || all_words[i].indexOf('?') != -1 || all_words[i].indexOf(';') != -1|| all_words[i].indexOf(')') != -1){
      temp_words.splice(t+1, 0, ".");
      temp_words.splice(t+1, 0, ".");
      temp_words.splice(t+1, 0, ".");
      t++;
      t++;
      t++;
    }

    t++;
  }
  all_words = temp_words.reverse();

  for (var i=0; i<all_words.length; i++){
    $(input_el).queue(showIt).delay(ms_per_word);
  }
}

function pauseReading(input_el) {
  paused_queue = $(input_el).queue();
  $(input_el).queue([]);
  $('#resume').removeClass('hidden');
  $('#pause').addClass('hidden');
}

function resumeReading(input_el) {
  $(input_el).queue(paused_queue);
  $(input_el).dequeue();
  $('#resume').addClass('hidden');
  $('#pause').removeClass('hidden');
}

function showIt(next) {
  var word = all_words.pop()
  document.getElementById('rsvp_result').innerHTML = word;
  next()
}
