package com.leeturner

import io.micronaut.configuration.picocli.PicocliRunner
import io.micronaut.context.ApplicationContext
import io.micronaut.context.env.Environment
import java.io.ByteArrayOutputStream
import java.io.PrintStream
import org.junit.jupiter.api.Test
import org.junit.jupiter.api.Assertions

class SnykMdCommandTest {

    @Test
    fun testWithCommandLineOption() {
        val ctx = ApplicationContext.run(Environment.CLI, Environment.TEST)
        val baos = ByteArrayOutputStream()
        System.setOut(PrintStream(baos))

        val args = arrayOf("-v")
        PicocliRunner.run(SnykMdCommand::class.java, ctx, *args)

        Assertions.assertTrue(baos.toString().contains("Hi!"))

        ctx.close()
    }
}
