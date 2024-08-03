# Go Blog

This project is a personal learning experience aimed at exploring new languages and technologies.

## Note

I have decided not to pursue writing my own Blog system in Go. While it is a cool little project to learn a language, it would not satisfy enough technological advances from my side. Even early during development, 
my first few thoughts were if instead of writing my backend i could just use an existing Headless CMS to save myself time and headaches. I then proceeded to write my Block Render Engine to at least get to know 
Go templates a bit better.

After having a kind of working tech demo, with constant Server Side Rendering of Blog articles, index pages, etc., I wanted to invest a bit of time into refactoring the code to improve the SSR performance. And what better
performance than to cache the rendered posts as static HTML files and include them. Then my thoughts went further, why not statically render _everything_ until a new blog post is created? While I have not proceeded
to implement a general static renderer, I simply could not see much more benefit to implement a tool that - in the end - would end up like Hugo, or other Static Site renderers. I was just rendering HTML with (at that point)
Lorem Ipsum structs. Nothing much more to learn from that point. 

The plan of distributing the blog via Kubernetes was another moonshot early on, as it would in the end just be a simple Docker image to build, and ship. I don't need to overarchitect this project. Even without the Static Rendering,
shipping the image to something like Google App Engine or Elastic Beanstalk would more than suffice for a site that maybe gets one visitor a week.
While, yes, I wanted to _learn_ the toolings, I simply could not do that to a satisfying degree with the project I've chosen. I will tackle this challenge with a distributed system that - while also probably not receiving any visitors - 
at least is a real world scenario for using the toolings I've chosen.

## Technologies

- Backend Language: Go (with Gorilla Mux)
- Frontend Technologies: HTMX, TailwindCSS, AlpineJS
- Data Storage: PostgreSQL
- Deployment Technologies: Docker, Kubernetes, Helm
- Dev Services: Devcontainers, Github Actions, act

## Features

- Learning new programming languages
- Exploring different technologies
- Hands-on experience with various development tasks
- Enhancing programming skills

## License

This project is licensed under the [MIT License](LICENSE).
