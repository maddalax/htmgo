## Getting Started


##### **Prerequisites:**
1. Go: https://go.dev/doc/install
2. Familiarity with https://htmx.org and html/hypermedia
   1. If you have not read the htmx docs, please do so before continuing, a lot of concepts below will be much more clear after.


<br>

##### 1. **Install htmgo**

```bash
GOPROXY=direct go install github.com/maddalax/htmgo/cli/htmgo@latest
```



**2. Create new project**
Once htmgo cli tool is installed, run

```bash
htmgo template
```

this will ask you for a new app name, and it will clone our starter template to a new directory it creates with your app name.

<br>

**3. Running the dev server**
htmgo has built in live reload on the dev server, to use this, run this command in the root of your project

```bash
htmgo watch
```

If you prefer to restart the dev server yourself (no live reload), use

```bash
htmgo run
```



##### **4. Core concepts**

View the [core concepts](/docs#core-concepts-pages) of how to use htmgo, such as adding pages, using partials, routing, etc.

<br>

**5. Building for production**
htmgo cli can be used to build the application for production as a single binary

```bash
htmgo build
```

it will be output to **./dist**



<br>
