$(function() {
  var term = $('body').terminal(function(command, term) {
    if (command === 'sections') {
      fetch('/api/sections')
        .then(response => resoponse.json())
        .then(data => {
          term.clear();
          data.forEach(section => {
            term.echo(`[[b;#f00;]${section.title}]`);
            term.echo(section.content);
            term.echo('');
          });
        })
        .catch(error => term.error(error));
    } else {
      term.echo('Invalid command. Type "sections" to view portfolio sections.')
    }
  }, {
    greetings: 'Welcome to my portfolio! Type "sections" to get started.',
    name: 'portfolio',
    height: 400,
    width: 800,
    prompt: '$ '
  });
});