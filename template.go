package main

var defaultTemplate = `<strong>{{ uppercasefirst build.status }}</strong> <a href=\"{{ system.link_url }}/{{ repo.owner }}/{{ repo.name }}/{{ build.number }}\">{{ repo.owner }}/{{ repo.name }}#{{ truncate build.commit 8 }}</a> ({{ build.branch }}) by {{ build.author }} in {{ duration build.started_at build.finished_at }} </br> - {{ build.message }}`
